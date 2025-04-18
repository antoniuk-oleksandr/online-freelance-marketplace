import type { MyProfileDeliveryRequestResponse } from "@/types/MyProfileDeliveryRequestResponse";

export const makeMyProfileDeliveryRequest = async (): Promise<MyProfileDeliveryRequestResponse> => {
  return {
    status: 200,
    data: {
      deliveryDate: 1742680800 * 1000,
      freelancer: {
        id: 1,
        username: 'alex',
      },
    }
  }
}
