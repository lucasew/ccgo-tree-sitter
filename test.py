# /// script
# requires-python = ">=3.11"
# dependencies = []
# ///

import os
import subprocess
import tempfile
import sys
import shutil
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
    # Create temp directory
    work_dir = tempfile.mkdtemp(prefix="tree-sitter-test-")
    # work_dir = "/tmp/lewtec_debug/tree-sitter-test"
    if os.path.exists(work_dir):
        shutil.rmtree(work_dir)
    os.makedirs(work_dir)
    print(f"Work dir: {work_dir}")

    try:
        # 1. Run Transpiler
        print("Running transpiler...")
        # Get absolute path to current directory (repo root)
        repo_root = os.getcwd()
        cmd = f"go run . -t /tmp/tree-sitter -g /tmp/tree-sitter-json -o {work_dir}"
        run_command(cmd, cwd=repo_root)

        # 2. Compile generated code
        print("Compiling...")
        run_command("go get modernc.org/libc@v1.68.0", cwd=work_dir)
        run_command("go mod tidy", cwd=work_dir)
        run_command("go build ./cmd/parse", cwd=work_dir)

        # 3. Check for suspicious padding patterns
        print("Checking for padding patterns in generated code...")
        core_file = os.path.join(work_dir, "core", "tree_sitter.go")

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

        # 4. Run Test
        print("Running test...")
        test_file = os.path.join(work_dir, "test.json")
        with open(test_file, "w") as f:
            f.write('{"key": [1, true]}')

        result = run_command("./parse test.json", cwd=work_dir)
        output = result.stdout.strip()
        print("Output:")
        print(output)

        # 5. Verify Output
        # Expected (partial):
        # document ...
        #   object ...
        #     pair ...
        #       key: ...
        #       value: ...

        checks = ["document [", "pair [", "key:", "value: array"]

        passed = True
        for check in checks:
            if check not in output:
                print(f"❌ Missing expected string: '{check}'")
                passed = False

        if passed:
            print("✅ Test PASSED")
        else:
            print("❌ Test FAILED: Output mismatch")
            sys.exit(1)

    except Exception as e:
        print(f"An error occurred: {e}")
        sys.exit(1)
    finally:
        # Cleanup (optional)
        # shutil.rmtree(work_dir)
        pass


if __name__ == "__main__":
    main()
