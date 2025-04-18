import { z } from "zod";

export const createServicePackagesFormSchema = z.object({
  files: z
    .array(z.instanceof(File))
    .min(1, "At least one picture is required"),
})
