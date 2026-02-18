# /// script
# requires-python = ">=3.11"
# dependencies = []
# ///

import os
import subprocess
import sys
import tempfile
import re


def run_command(command, cwd=None, check=True):
    """Run a shell command and return its output."""
    try:
        result = subprocess.run(
            command,
            cwd=cwd,
            check=check,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True,
            shell=True,
        )
        return result
    except subprocess.CalledProcessError as e:
        print(f"Command failed: {command}")
        print("STDOUT:", e.stdout)
        print("STDERR:", e.stderr)
        raise


def main():
    repo_root = os.getcwd()

    try:
        # 1. Run Codegen
        print("Running codegen...")
        run_command("go run ./cmd/codegen", cwd=repo_root)

        # 3. Check for suspicious padding patterns in core
        print("Checking for padding patterns in generated core...")
        # Find the core file (it has GOOS and GOARCH in the name)
        grammar_dir = os.path.join(repo_root, "grammar")
        core_files = [
            f
            for f in os.listdir(grammar_dir)
            if f.startswith("core-") and f.endswith(".go")
        ]

        if not core_files:
            print("❌ Could not find generated core file in grammar/")
            sys.exit(1)

        core_file = os.path.join(grammar_dir, core_files[0])
        print(f"Analyzing {core_file}...")

        with open(core_file, "r") as f:
            content = f.read()

        # Pattern: &struct { _ [N]byte
        if re.search(r"&struct \{ _ \[", content):
            print("⚠️  Found struct padding in generated code!")
            # Print context (grep-like)
            lines = content.split("\n")
            for i, line in enumerate(lines):
                if "&struct { _ [" in line:
                    print(f"Line {i + 1}: {line.strip()}")
                    # Print context
                    start = max(0, i - 2)
                    end = min(len(lines), i + 3)
                    for ctx_line in lines[start:end]:
                        print(f"  {ctx_line.strip()}")
                    break  # Just show first one
        else:
            print("✅ No struct padding found in generated code (or N=0)")

        # 4. Run Test with JSON
        print("Running test with JSON...")
        with tempfile.NamedTemporaryFile(suffix=".json", mode="w", delete=False) as f:
            f.write('{"key": [1, true]}')
            test_json = f.name

        try:
            result = run_command(f"go run ./cmd/parse {test_json}", cwd=repo_root)
            output = result.stdout.strip()
            print("Output:")
            print(output)

            # 5. Verify Output
            checks = ["document [", "pair [", "key:", "value: array"]
            passed = True
            for check in checks:
                if check not in output:
                    print(f"❌ Missing expected string: '{check}'")
                    passed = False

            if passed:
                print("✅ JSON Test PASSED")
            else:
                print("❌ JSON Test FAILED: Output mismatch")
                sys.exit(1)
        finally:
            if os.path.exists(test_json):
                os.remove(test_json)

        # 6. Run Test with Lua (if available)
        if os.path.exists(os.path.join(repo_root, "grammar/lua")):
            print("Running test with Lua...")
            with tempfile.NamedTemporaryFile(
                suffix=".lua", mode="w", delete=False
            ) as f:
                f.write("local x = 10\nprint(x)")
                test_lua = f.name
            try:
                result = run_command(f"go run ./cmd/parse {test_lua}", cwd=repo_root)
                print("Lua Output:")
                print(result.stdout.strip())
                if "chunk [" in result.stdout:
                    print("✅ Lua Test PASSED")
                else:
                    print("❌ Lua Test FAILED")
                    passed = False
            finally:
                if os.path.exists(test_lua):
                    os.remove(test_lua)

    except Exception as e:
        print(f"An error occurred: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
