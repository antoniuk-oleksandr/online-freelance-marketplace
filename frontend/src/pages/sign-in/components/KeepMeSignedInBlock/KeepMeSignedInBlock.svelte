<script lang="ts">
  import KeepMeSignedInBlockLayout from '@/pages/sign-in/components/KeepMeSignedInBlock/KeepMeSignedInBlockLayout.svelte'
  import KeepMeSignedInBlockInput from '@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockInput/KeepMeSignedInBlockInput.svelte'
  import KeepMeSignedInBlockText from '@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockText/KeepMeSignedInBlockText.svelte'
  import Checkbox from '@/common-components/Checkbox/Checkbox.svelte'

  type KeepMeSignedInBlockProps = {
    keepSignedIn: boolean
  }

  let { keepSignedIn = $bindable() }: KeepMeSignedInBlockProps = $props()
  const setKeepSignedIn = (value: boolean) => (keepSignedIn = value)
</script>

<KeepMeSignedInBlockLayout {keepSignedIn} {setKeepSignedIn}>
  <Checkbox selected={keepSignedIn} />
  <KeepMeSignedInBlockInput {keepSignedIn} />
  <KeepMeSignedInBlockText />
</KeepMeSignedInBlockLayout>
