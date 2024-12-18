import {Language} from "@/types/Language.ts";
import {Skill} from "@/types/Skill.ts";
import {Category} from "@/types/Category.ts";
import {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";

export type SearchFilterArrayStore = {
    data: Language[] | Skill[] | Category[],
    isOpened: boolean,
    title: string,
    attribute: SearchFilterArrayAttribute,
}