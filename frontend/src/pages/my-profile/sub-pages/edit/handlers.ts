import type { MyProfileEditPageFormData } from "@/types/MyProfileEditPageFormData";
import { findLanguageOrSkill, getMyProfileEditPageDataDifference } from "./helpers";
import { modalStore } from "@/common-stores/modal-store";
//@ts-ignore
import { myProfileEditPageLanguagesSkillsModal } from "./components/MyProfileEditPageLanguagesSkillsModal/MyProfileEditPageLanguagesSkillsModal.svelte"
import type { FilterParams } from "@/types/FilterParams";
import type { Writable } from "svelte/store";

export const handleMyProfileEditPageFormSubmit = async (
  data: MyProfileEditPageFormData,
  initialData: MyProfileEditPageFormData,
) => {
  const diff = getMyProfileEditPageDataDifference(data, initialData);
  console.log("Data difference:", diff);
}

export const handleAddLanguagesSkillsButtonClick = (
  formDataStore: Writable<MyProfileEditPageFormData>,
  filterParams: FilterParams,
  type: 'languages' | 'skills',
) => {
  modalStore.set({
    title: "Add a " + type.slice(0, -1),
    isOpened: true,
    renderContent: myProfileEditPageLanguagesSkillsModal,
    props: {
      type,
      formDataStore,
      filterParams,
    },
  })
}

export const handleLanguagesSkillsModalSubmit = (
  selectedValue: string,
  type: 'languages' | 'skills',
  formDataStore: Writable<MyProfileEditPageFormData>,
  filterParams: FilterParams
) => {
  const item = findLanguageOrSkill(filterParams, type, selectedValue);

  formDataStore.update((prev) => {
    return ({ ...prev, [type]: [...prev[type], item] })
  })

  modalStore.set({ isOpened: false })
}
