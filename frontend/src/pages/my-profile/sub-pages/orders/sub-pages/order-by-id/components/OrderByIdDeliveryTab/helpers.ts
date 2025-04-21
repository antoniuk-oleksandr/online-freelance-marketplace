import type { MyProfileDeliveryRequestResponse } from "@/types/MyProfileDeliveryRequestResponse";

export const makeMyProfileDeliveryRequest = async (): Promise<MyProfileDeliveryRequestResponse> => {
  return {
    status: 200,
    data: {
      deliveryDate: 1741118830000,
      freelancer: {
        id: 1,
        username: 'name',
      },
    }
  }
}
