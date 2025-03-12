<script lang="ts">
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
    error: string | undefined
    styles?: string
  }

  let focus = $state(false)
  let setFocus = (value: boolean) => (focus = value)

  const props: TextareaProps = $props()
</script>

<TextareaLayout>
  <TextareaContainerLayout {focus} {...props}>
    <TextareaElement {setFocus} {...props} />
    <TextareaCounter {focus} {...props}/>
  </TextareaContainerLayout>
  <InputError {...props} />
</TextareaLayout>
