import type { Category } from "./Category";
import type { Language } from "./Language";
import type { Skill } from "./Skill";

export type FilterParams = {
  skill: Skill[],
  category: Category[],
  language: Language[],
}
