<script lang="ts">
  import type { SignProps } from '@/types/SignProps.ts'
  import Button from '@/common-components/Button/Button.svelte'
  import LeftSignSideOrText from '@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrText.svelte'
  import SignFormLayout from '@/common-components/Sign/components/SignForm/SignFormLayout.svelte'
  import GoogleButton from '@/common-components/Sign/components/GoogleButton/GoogleButton.svelte'
  import BackSignInButton from '@/common-components/Sign/components/BackSignInButton/BackSignInButton.svelte'

  type SignFormProps = SignProps & {
    keepSignedIn: boolean
  }

  let {
    keepSignedIn = $bindable(),
    showBackButton,
    children,
    signButtonText,
    ...rest
  }: SignFormProps = $props()
  let loading = $state(false)
  const setLoading = (value: boolean) => (loading = value)
</script>

<SignFormLayout bind:keepSignedIn {setLoading} {signButtonText} {...rest}>
  {#if children}
    {@render children()}
  {/if}
  <Button {loading} styles="text-lg !h-12" type={'submit'}>{signButtonText}</Button>
  {#if showBackButton}
    <BackSignInButton />
  {/if}
  <LeftSignSideOrText {signButtonText} {...rest} />
  <GoogleButton {keepSignedIn} {signButtonText} {...rest} />
</SignFormLayout>
