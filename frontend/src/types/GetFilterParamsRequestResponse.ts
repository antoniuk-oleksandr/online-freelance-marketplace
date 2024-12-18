import {Skill} from "@/types/Skill.ts";
import {Category} from "@/types/Category.ts";
import {Language} from "@/types/Language.ts";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

export type GetFilterParamsRequestResponse = ResponseError & {
    status: number,
    data?: {
        skill: Skill[],
        category: Category[],
        language: Language[],
    },
}