import type { Language } from "./Language";
import type { Skill } from "./Skill";

export type MyProfileEditPageFormData = {
  email: string,
  username: string,
  firstName: string,
  surname: string,
  about: string,
  languages: Language[],
  skills: Skill[],
  avatar: File | null | string,
}
