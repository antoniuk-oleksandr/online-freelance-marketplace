<script lang="ts">
  import type { SignProps } from '@/types/SignProps.ts'
  import SignLayout from '@/common-components/Sign/SignLayout.svelte'
  import LeftSignSide from '@/common-components/Sign/components/LeftSignSide/LeftSignSide.svelte'
  import RightSignSide from '@/common-components/Sign/components/RightSignSide/RightSignSide.svelte'
  import SentEmailMessage from '@/common-components/Sign/components/SentEmailMessage/SentEmailMessage.svelte'

  type SignElementProps = SignProps & {
    keepSignedIn: boolean
  }

  let { keepSignedIn = $bindable(), showEmailSentMessage, ...rest }: SignElementProps = $props()
</script>

{#if showEmailSentMessage}
  <SentEmailMessage />
{:else}
  <SignLayout>
    <LeftSignSide bind:keepSignedIn {...rest} />
    <RightSignSide />
  </SignLayout>
{/if}
