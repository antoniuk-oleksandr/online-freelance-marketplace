import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { FilterParams } from "@/types/FilterParams";
import type { GetFilterParamsRequestResponse } from "@/types/GetFilterParamsRequestResponse";
import type { Language } from "@/types/Language";
import type { MyProfileDataRequestResponse } from "@/types/MyProfileDataRequestResponse";
import type { MyProfileEditPageDifferentData } from "@/types/MyProfileEditPageDifferentData";
import type { MyProfileEditPageFormData } from "@/types/MyProfileEditPageFormData";
import type { MyProfileEditPageInputData } from "@/types/MyProfileEditPageInputData";
import type { NewSelectItem } from "@/types/NewSelectItem";
import type { Skill } from "@/types/Skill";
import { z } from "zod";

export const myProfileEditPageFormScheme = z.object({
  email: z.string().min(1, { message: "Email is required." }),
  username: z.string().min(1, { message: "Username is required." }),
  firstName: z.string().min(1, { message: "First name is required." }),
  surname: z.string().min(1, { message: "Surname is required." }),
  about: z.string().optional(),
  languages: z.array(z.object({
    id: z.number(),
    name: z.string(),
  })).optional(),
  skills: z.array(z.object({
    id: z.number(),
    name: z.string(),
  })).optional(),
  avatar: z.any().optional(),
})

export const makeMyProfileDataRequest = async (): Promise<MyProfileEditPageFormData | null> => {
  const response = {
    status: 200,
    data: {
      email: "jane.doe@example.com",
      username: "janedoe",
      firstName: "Jane",
      surname: "Doe",
      about: "<p>Passionate frontend developer with a love for clean UI and smooth UX.</p>",
      languages: [{ name: "English", id: 1 }, { name: "Mandarin Chinese", id: 2 }],
      skills: [{ name: "HTML & CSS", id: 1 }],
      avatar: "http://localhost:8030/files/image_1.jpg",
    }
  } as MyProfileDataRequestResponse;


  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error })
    return null
  }

  return response.data;
}

export const makeFilerParamsRequest = async (): Promise<FilterParams | null> => {
  const response = await request<GetFilterParamsRequestResponse>('GET', '/filter-params/all')

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error })
    return null
  }

  return response.data;
}

export const fetchProfileDataAndFilters = async (
): Promise<[MyProfileEditPageFormData | null, FilterParams | null]> => Promise.all([
  makeMyProfileDataRequest(),
  makeFilerParamsRequest(),
])

export const getMyProfileEditPageInputsData = (
  formErrors: Record<string, string[]>
): MyProfileEditPageInputData[] => [{
  error: formErrors.email && formErrors.email[0],
  id: "email",
  label: "Email",
  placeholder: "Enter your email",
},
{
  error: formErrors.username && formErrors.username[0],
  id: "username",
  label: "Username",
  placeholder: "Enter your username",
},
{
  error: formErrors.firstName && formErrors.firstName[0],
  id: "firstName",
  label: "First Name",
  placeholder: "Enter your first name",
},
{
  error: formErrors.surname && formErrors.surname[0],
  id: "surname",
  label: "Surname",
  placeholder: "Enter your surname",
}]

export const getMyProfileEditPageDataDifference = (
  data: MyProfileEditPageFormData,
  initialData: MyProfileEditPageFormData,
): MyProfileEditPageDifferentData => {
  const keys: (keyof MyProfileEditPageFormData)[] = ["email", "username", "firstName", "surname", "about", "avatar"];

  const dataDifference: MyProfileEditPageDifferentData = {}

  keys.forEach((key) => {
    if (data[key] !== initialData[key]) {
      //@ts-ignore
      dataDifference[key] = data[key]
    } else {
      //@ts-ignore
      dataDifference[key] = null
    }
  })

  const languageDifference = getDifference(data.languages, initialData.languages);
  dataDifference.languagesToAdd = languageDifference.toAdd;
  dataDifference.languagesToRemove = languageDifference.toRemove;

  const skillDifference = getDifference(data.skills, initialData.skills);
  dataDifference.skillsToAdd = skillDifference.toAdd;
  dataDifference.skillsToRemove = skillDifference.toRemove;

  return dataDifference
}

const getDifference = <T extends { id: number }>(
  current: T[] | undefined,
  initial: T[] | undefined
) => ({
  toAdd: current?.filter((item) => !initial?.some((initialItem) => initialItem.id === item.id)).map((item) => item.id),
  toRemove: initial?.filter((item) => !current?.some((currentItem) => currentItem.id === item.id)).map((item) => item.id),
});

export const extractNewSelectItemsFromFilterParams = (
  filterParams: FilterParams,
  type: 'languages' | 'skills',
  formData: MyProfileEditPageFormData
): NewSelectItem[] => {
  return filterParams[type.slice(0, -1) as keyof FilterParams]
    .reduce((acc: NewSelectItem[], item) => {
      if (formData[type].some((formItem) => formItem.id === item.id)) return acc
      return acc.concat({ title: item.name, })
    }, [])
}

export const findLanguageOrSkill = (
  filterParams: FilterParams, type: 'languages' | 'skills', value: string
): Language | Skill =>
  filterParams[type.slice(0, -1) as keyof FilterParams]
    .find((item) => item.name === value) as Language | Skill

