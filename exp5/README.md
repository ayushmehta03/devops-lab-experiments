#git reset and revert

git reset-> used to move the current branch head to a specific previous commit.

types:

Soft Reset (--soft): Moves the HEAD to the specified commit but keeps all your changes in the staging area

Mixed Reset (Default): Moves the HEAD to the specified commit and resets the staging area, but keeps your changes in the working directory as unstaged.

Hard Reset (--hard): Moves the HEAD to the specified commit and discards all changes in both the staging area and working directory.


git revert-> The git revert command is a "forward-moving" undo operation. Instead of deleting history, it creates a new commit that introduces the exact opposite changes of the specified commit. 



