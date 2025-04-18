import type { NewSelectItem } from "@/types/NewSelectItem"
import type { StepperItem } from "@/types/StepperItem"
import { getServiceCategories } from "@/utils/utils"
import { z } from "zod"

export const getCategorySelectInitialItems = (): NewSelectItem[] => {
  const categories = getServiceCategories()

  return categories.map((item) => ({
    title: item.name,
  }))
}

export const createServiceBasicInfoFormSchema = z.object({
  title: z.string().min(1, { message: "Title is required." }),
  description: z.
    string().
    refine((value) => !isRichTextEmpty(value), { message: "Description is required." }),
  categoryId: z.number().min(1, { message: "Category is required." }),
})

function isRichTextEmpty(html: string): boolean {
  const div = document.createElement('div');
  div.innerHTML = html;

  const text = div.textContent?.trim();

  return !text;
}
