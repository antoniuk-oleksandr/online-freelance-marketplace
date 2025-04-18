import {z} from "zod";

export const initialForgotPasswordData: any = {
    email: "",
}

export const forgotPasswordSchema = z.object({
    email: z.string().min(1, {message: "Email address is required."}),
})
