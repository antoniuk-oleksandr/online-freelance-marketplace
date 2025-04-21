<script lang="ts">
  import LabeledInputError from './components/LabeledInputError/LabeledInputError.svelte'
  import LabeledInputElement from './components/LabeledInputElement/LabeledInputElement.svelte'
  import LabeledInputElementLabel from './components/LabeledInputElementLabel/LabeledInputElementLabel.svelte'
  import LabeledInputLayout from './LabeledInputLayout.svelte'

  type CardInputElementProps = {
    label: string
    id: string
    placeholder: string
    type?: string
    error?: string | null
    styles?: string
    focusAction?: () => void
    blurAction?: () => void
  }

  let props: CardInputElementProps = $props()
</script>

<LabeledInputLayout styles={props.styles}>
  <LabeledInputElementLabel {...props} />
  <LabeledInputElement {...props} />
  <LabeledInputError {...props} />
</LabeledInputLayout>
