import { Skill } from "@/types/Skill";
import { Category } from "@/types/Category";
import { Language } from "@/types/Language";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";

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
