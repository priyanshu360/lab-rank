import sys
import json

def run_test(test_input, expected_output, function_name):
    try:
        # Execute the solution code in the global scope
        globals_dict = {}
        exec(solution_code, globals_dict)
        if function_name not in globals_dict or not callable(globals_dict[function_name]):
            return {"test_case": test_input, "error": "Error: The function {} is not defined.".format(function_name), "result": "Failed"}

        # Test the function with the provided input
        result = globals_dict[function_name](test_input)

        # Check if the result matches the expected output
        if result != expected_output:
            return {"test_case": test_input, "error": None, "result": "Failed", "details": {"expected_output": expected_output, "actual_output": result}}

        # If the result matches the expected output, return success
        return {"test_case": test_input, "error": None, "result": "Passed"}

    except Exception as e:
        return {"test_case": test_input, "error": "Error: {}".format(str(e)), "result": "Failed"}

def check_fibonacci_solution(solution_file_path, test_cases):
    try:
        with open(solution_file_path, 'r') as file:
            global solution_code
            solution_code = file.read()

        results = []
        for test_case in test_cases:
            result = run_test(test_case["input"], test_case["output"], test_case["function_name"])
            results.append(result)

        return json.dumps(results, indent=2)

    except Exception as e:
        return json.dumps([{"error": "Error: {}".format(str(e))}], indent=2)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python check_solution.py <solution_file_path>")
        sys.exit(1)

    # Define 10 test cases with input, expected output, and function name
    test_cases = [
        {"input": 0, "output": [], "function_name": "generate_fibonacci"},
        {"input": 1, "output": [0], "function_name": "generate_fibonacci"},
        {"input": 2, "output": [0, 1], "function_name": "generate_fibonacci"},
        {"input": 5, "output": [0, 1, 1, 2, 3], "function_name": "generate_fibonacci"},
        {"input": 10, "output": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34], "function_name": "generate_fibonacci"},
        {"input": 15, "output": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377], "function_name": "generate_fibonacci"},
        {"input": 20, "output": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181], "function_name": "generate_fibonacci"},
        {"input": 25, "output": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368], "function_name": "generate_fibonacci"},
        {"input": 30, "output": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229], "function_name": "generate_fibonacci"},
    ]

    solution_file_path = sys.argv[1]
    result = check_fibonacci_solution(solution_file_path, test_cases)
    print(result)
