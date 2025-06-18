import type { DeliveryCancellation } from "./DeliveryCancellation";
import type { DeliveryDataFreelancer } from "./DeliveryDataFreelancer";
import type { DeliveryDataInfo } from "./DeliveryDataInfo";
import type { DeliveryPayment } from "./DeliveryPayment";
import type { StatusEnum } from "./StatusEnum";

export type DeliveryData = {
  status: StatusEnum,
  delivery: DeliveryDataInfo | null,
  freelancer: DeliveryDataFreelancer | null,
  payment: DeliveryPayment | null,
  cancellation: DeliveryCancellation | null,
} 
