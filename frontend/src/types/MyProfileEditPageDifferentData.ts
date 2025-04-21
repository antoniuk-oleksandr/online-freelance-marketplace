export type MyProfileEditPageDifferentData = {
  email?: string,
  username?: string,
  firstName?: string,
  surname?: string,
  about?: string,
  avatar?: File | null | string,
  languagesToAdd?: number[],
  languagesToRemove?: number[],
  skillsToAdd?: number[],
  skillsToRemove?: number[],
}
