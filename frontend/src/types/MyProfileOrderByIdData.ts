import type { MyProfileChatRequestResponse } from "./MyProfileChatRequestResponse";
import type { MyProfileDeliveryRequestResponse } from "./MyProfileDeliveryRequestResponse";
import type { MyProfileDiaryRequestResponse } from "./MyProfileDiaryRequestResponse";
import type { MyProfileOverviewRequestResponse } from "./MyProfileOverviewRequestResponse";
import type { MyProfileRequirementsRequestResponse } from "./MyProfileRequirementsRequestResponse";

export type MyProfileOrderByIdData = [
  MyProfileOverviewRequestResponse | undefined,
  MyProfileRequirementsRequestResponse | undefined,
  MyProfileChatRequestResponse | undefined,
  MyProfileDiaryRequestResponse | undefined,
  MyProfileDeliveryRequestResponse | undefined,
];

