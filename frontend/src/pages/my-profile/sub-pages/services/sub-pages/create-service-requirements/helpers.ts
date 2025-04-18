import { modalStore } from "@/common-stores/modal-store";
import type { CreateServiceRequirementsFormData } from "@/types/CreateServiceRequirementsFormData";
import type { Writable } from "svelte/store";
import { z } from "zod";
//@ts-ignore
import { CreateServiceQuestionModal } from './components/CreateServiceQuestionModal/CreateServiceQuestionModal.svelte'

export const createServiceRequirementsFormSchema = z.object({
  questions: z.array(z.object({
    text: z.string().min(1, "Question is required."),
    id: z.number(),
  })).min(1, "At least one question is required."),
})

export const addQuestionSubmitAction = (
  value: string,
  formDataStore: Writable<CreateServiceRequirementsFormData>
) => {
  formDataStore.update((prev) => {
    const newQuestion = {
      text: value,
      id: prev.questions.length,
    }

    return {
      ...prev,
      questions: [...prev.questions, newQuestion],
    }
  })
  modalStore.set({ isOpened: false })
}

export const renderAddQuestionModal = (
  formDataStore: Writable<CreateServiceRequirementsFormData>
) => {
  modalStore.set({
    isOpened: true,
    renderContent: CreateServiceQuestionModal,
    props: {
      submitAction: (value: string) => addQuestionSubmitAction(value, formDataStore),
    },
    title: 'Add a question',
  })
}
