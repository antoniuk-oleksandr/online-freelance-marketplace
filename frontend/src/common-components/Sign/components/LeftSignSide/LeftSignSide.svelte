<script lang="ts">
  import LeftSignSideLayout from '@/common-components/Sign/components/LeftSignSide/LeftSignSideLayout.svelte'
  import LeftSignSideTopText from '@/common-components/Sign/components/LeftSignSideTopText/LeftSignSideTopText.svelte'
  import SignForm from '@/common-components/Sign/components/SignForm/SignForm.svelte'
  import SubFormText from '@/common-components/Sign/components/SubFormText/SubFormText.svelte'
  import type { SignProps } from '@/types/SignProps'

  type LeftSignSideProps = SignProps & {
    keepSignedIn: boolean
  }

  let { keepSignedIn = $bindable(), ...rest }: LeftSignSideProps = $props()
</script>

<LeftSignSideLayout>
  <LeftSignSideTopText {...rest} />
  <SignForm bind:keepSignedIn {...rest} />
  <SubFormText {...rest as any} />
</LeftSignSideLayout>
