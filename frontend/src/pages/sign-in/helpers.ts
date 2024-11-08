import {z, ZodObject} from 'zod';

export const signInSchema: ZodObject<any> = z.object({
    email: z.string().min(1, "Email is required"),
    password: z.string().min(1, "Password is required"),
});

export const defaultSignInValues = {
    email: "",
    password: "",
};
