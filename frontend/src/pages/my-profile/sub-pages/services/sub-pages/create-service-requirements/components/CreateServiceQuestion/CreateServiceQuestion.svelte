<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import Icon from '@iconify/svelte'
  import CreateServiceQuestionLayout from './CreateServiceQuestionLayout.svelte'

  type CreateServiceQuestionProps = {
    id: number
    question: string
    removeQuestion: (index: number) => void
  }

  const { id, question, removeQuestion }: CreateServiceQuestionProps = $props()
</script>

<CreateServiceQuestionLayout>
  <div>
    <Label text="Question {id + 1}" />
    <p>{question}</p>
  </div>
  <div
    class="opacity-0 group-hover:opacity-100 right-0 top-0 absolute -mr-[17px] -mt-[17px] bg-light-palette-background-block dark:bg-dark-palette-background-block rounded-full p-2 text-red-500 hover:bg-red-100 dark:hover:bg-red-500/10 hover:text-red-600 transition-all ease-in-out duration-200 cursor-pointer focus:outline-none focus:ring-2 focus:ring-red-500"
  >
    <Icon onclick={() => removeQuestion(id)} width={18} height={18} icon="ci:close-lg" />
  </div>
</CreateServiceQuestionLayout>
