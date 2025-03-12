<script lang="ts">
  import TextArea from '@/common-components/Textarea/Textarea.svelte'
  import OrderSubmitRequirementsFormQuestionLayout from './OrderSubmitRequirementsFormQuestionLayout.svelte'

  type OrderSubmitRequirementsFormQuestionProps = {
    content: string
    id: number | string
    error: string | undefined
    placeholder: string
    index?: number
    value: string
    wasSubmitted: boolean
  }

  const { content, index, id, ...rest }: OrderSubmitRequirementsFormQuestionProps = $props()
</script>

<OrderSubmitRequirementsFormQuestionLayout>
  <span>{index !== undefined ? index + 1 + '.' : ''} {content}</span>
  <TextArea useCounter  id={id.toString()} name={id.toString()} {...rest} />
</OrderSubmitRequirementsFormQuestionLayout>
