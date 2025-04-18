<script module lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import Textarea from '@/common-components/Textarea/Textarea.svelte'
  import type { effect } from 'zod'

  type CreateServiceQuestionModalProps = {
    submitAction: (value: string) => void
  }

  let question = $state('')
  const setValue = (value: string) => (question = value)

  export { CreateServiceQuestionModal }
</script>

{#snippet CreateServiceQuestionModal(props: CreateServiceQuestionModalProps)}
  <div class="p-6 flex gap-6 flex-col">
    <Textarea
      {setValue}
      value={question}
      wasSubmitted={false}
      placeholder="Type here"
      id="question"
      name="question"
      useCounter
    />
    <Button
      clickAction={() => {
        props.submitAction(question)
        setValue('')
      }}
      styles="!w-40 self-end"
      type="button">Submit</Button
    >
  </div>
{/snippet}
