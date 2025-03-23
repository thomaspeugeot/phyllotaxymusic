#!/usr/bin/env bash

# Loop through all files in the current directory matching the prefix "parameter_"
for old_name in parameter_update*; do

  # Ensure that the file actually exists (i.e., that the pattern matched at least one file)
  if [ -f "$old_name" ]; then
    
    # Build the new name by stripping "parameter_" and prepending "stager_"
    new_name="stager_update${old_name#parameter_update}"

    # Rename the file
    mv "$old_name" "$new_name"
    echo "Renamed: $old_name -> $new_name"

  fi
done

echo "File renaming completed!"
