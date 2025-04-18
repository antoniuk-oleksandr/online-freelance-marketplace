<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Icon from '@iconify/svelte'
  import CreateServiceQuestionList from '../CreateServiceQuestionList/CreateServiceQuestionList.svelte'
  import type { CreateServiceRequirementsFormData } from '@/types/CreateServiceRequirementsFormData'
  import type { Writable } from 'svelte/store'
  import { renderAddQuestionModal } from '../../helpers'
  import CreateServiceQuestionsBlockLayout from './CreateServiceQuestionsBlockLayout.svelte'
  import CreateServiceQuestionMesage from '../CreateServiceQuestionMesage/CreateServiceQuestionMesage.svelte'
  import LabeledInputError from '@/common-components/LabeledInput/components/LabeledInputError/LabeledInputError.svelte'

  type CreateServiceQuestionsBlockProps = {
    formDataStore: Writable<CreateServiceRequirementsFormData>
    formErrorsStore: Writable<Record<string, string[]>>
  }
  const { formDataStore, formErrorsStore }: CreateServiceQuestionsBlockProps = $props()

  const setQuestions = (value: { text: string; id: number }[]) => {
    formDataStore.update((prev) => ({ ...prev, questions: value }))
  }

  const removeQuestion = (index: number) => {
    formDataStore.update((prev) => ({
      ...prev,
      questions: prev.questions.filter((item) => item.id !== index),
    }))
  }
</script>

<CreateServiceQuestionsBlockLayout>
  <CreateServiceQuestionMesage />
  <CreateServiceQuestionList {removeQuestion} questions={$formDataStore.questions} {setQuestions} />
  <Button
    clickAction={() => renderAddQuestionModal(formDataStore)}
    styles="w-fit flex gap-1 items-center mt-6"
    color="grey"
  >
    <div class="size-6">
      <Icon width={24} height={24} icon="hugeicons:plus-sign" />
    </div>
    <span>Add a question</span>
  </Button>
  <LabeledInputError styles="mt-1" error={$formErrorsStore.questions && $formErrorsStore.questions[0]} />
</CreateServiceQuestionsBlockLayout>
