# -depth ensures that find processes a directory's contents before the directory itself
# This helps avoid issues when renaming directories that are still being traversed.
find . -depth -name '*phyllotaxymusic*' | while IFS= read -r old_path; do
  # Generate the new path by replacing the substring
  new_path="${old_path//phyllotaxymusic/phyllotaxymusic}"

  echo "Renaming:"
  echo "  $old_path"
  echo "  -> $new_path"
  echo

  # Perform the move (rename)
  mv -- "$old_path" "$new_path"
done

echo "Done renaming."