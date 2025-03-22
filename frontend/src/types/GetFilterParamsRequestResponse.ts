import { Skill } from "@/types/Skill";
import { Category } from "@/types/Category";
import { Language } from "@/types/Language";
import { type ResponseError } from "@/types/ResponseErrorEnum";

export type GetFilterParamsRequestResponse = {
    data: {
        skill: Skill[],
        category: Category[],
        language: Language[],
    },
    status: 200,
} | {
    data: ResponseError,
    status: 404 | 500 | 401,
}

