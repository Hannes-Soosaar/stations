
# Test the two-four.map
cd cmd && go run .  two-four.map two four 5

output=$(response)
echo "The output is: $output"

if [[ "$output" == *"expected_string"* ]]; then
    echo "The output contains the expected string."
else
    echo "The output does not contain the expected string."
fi
