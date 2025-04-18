import { z } from "zod";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import { errorStore } from "@/common-stores/error-store";

export const initialResetPasswordData = {
    password: '',
    confirmPassword: ''
}

export const resetPasswordSchema = z.object({
    password: z.string()
        .min(1, { message: "Password is required." })
        .min(8, { message: "Password must be at least 8 characters long." })
        .regex(/[a-z]/, { message: "Password must contain at least one lowercase letter." })
        .regex(/[A-Z]/, { message: "Password must contain at least one uppercase letter." })
        .regex(/[0-9]/, { message: "Password must contain at least one digit." })
        .regex(/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+/, { message: "Password must contain at least one special character." }),
    confirmPassword: z.string()
        .min(1, { message: "Confirm password is required." })
}).refine(data => data.password === data.confirmPassword, {
    message: "Passwords do not match.",
    path: ['confirmPassword']
})

export const processUrlToken = (
    setToken: (token: string) => void,
) => {
    const params = new URLSearchParams(window.location.search);
    const tokenParam = params.get('token') ?? "";
    setToken(tokenParam);

    if (!tokenParam) {
        errorStore.set({ shown: true, error: ResponseErrorEnum.InvalidToken });
        return;
    }
}
