<script lang="ts">
  import Label from '../Label/Label.svelte'
  import InputError from '../Sign/components/SignInput/components/InputError/InputError.svelte'
  import TextareaContainerLayout from './components/TextareaContainer/TextareaContainerLayout.svelte'
  import TextareaCounter from './components/TextareaCounter/TextareaCounter.svelte'
  import TextareaElement from './components/TextareaElement/TextareaElement.svelte'
  import TextareaLayout from './TextareaLayout.svelte'

  type TextareaProps = {
    placeholder: string
    id: string
    name: string
    wasSubmitted: boolean
    useCounter: boolean
    value?: string
    label?: string
    error?: string
    styles?: string
    setValue?: (value: string) => void
  }

  let focus = $state(false)
  let setFocus = (value: boolean) => (focus = value)

  let props: TextareaProps = $props()
</script>

<TextareaLayout>
  {#if props.label}
    <Label text={props.label} />
  {/if}
  <TextareaContainerLayout {focus} {...props}>
    <TextareaElement {setFocus} {...props} />
    <TextareaCounter {focus} {...props} />
  </TextareaContainerLayout>
  {#if props.error}
    <InputError error={props.error} {...props} />
  {/if}
</TextareaLayout>
