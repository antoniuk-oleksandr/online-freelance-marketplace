import {z, ZodObject} from "zod";

export const signUpSchema: ZodObject<any> = z.object({
    name: z.string().min(1, "Name is required").min(3, "Name must be at least 3 characters long"),
    email: z.string().min(1, "Email is required").email(),
    password: z
        .string()
        .min(1, "Password is required")
        .min(8, "Password must be at least 8 characters long")
        .regex(/[a-z]/, "Password must contain at least one lowercase letter")
        .regex(/[A-Z]/, "Password must contain at least one uppercase letter")
        .regex(/\d/, "Password must contain at least one number")
        .regex(/[@$!%*?&]/, "Password must contain at least one special character")

});

export const defaultSignUpValues = {
    name: "",
    email: "",
    password: "",
}