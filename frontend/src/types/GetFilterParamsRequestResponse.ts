import {Skill} from "@/types/Skill.ts";
import {Category} from "@/types/Category.ts";
import {Language} from "@/types/Language.ts";
import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export type GetFilterParamsRequestResponse = {
    data: {
        skill: Skill[],
        category: Category[],
        language: Language[],
    },
    status: 200,
} | {
    error: ResponseErrorEnum,
    status: 404 | 500,
}