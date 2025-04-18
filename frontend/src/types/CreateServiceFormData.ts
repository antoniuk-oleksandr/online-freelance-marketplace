import type { CreateServiceBasicInfoFormData } from "./CreateServiceBasicInfoFormData";
import type { CreateServiceMediaFormData } from "./CreateServiceMediaFormData";
import type { CreateServicePackageFormData } from "./CreateServicePackageFormData";
import type { CreateServiceRequirementsFormData } from "./CreateServiceRequirementsFormData";

export type CreateServiceFormData =
  CreateServiceBasicInfoFormData
  & CreateServicePackageFormData
  & CreateServiceRequirementsFormData
  & CreateServiceMediaFormData
