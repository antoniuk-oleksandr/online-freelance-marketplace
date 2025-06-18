import type { DeliveryData } from "@/types/DeliveryData";
import { z } from "zod";

export const deliveryFormSchema = z.object({
  files: z.array(z.instanceof(File)).min(1, { message: "At least one file is required." }),
  message: z.string().min(1, { message: "Message is required." }),
})

export const compareDeliveryData = (
  oldDeliveryData: DeliveryData,
  newDeliveryData: DeliveryData | undefined
): DeliveryData => {
  return !newDeliveryData ? oldDeliveryData : newDeliveryData;
}
