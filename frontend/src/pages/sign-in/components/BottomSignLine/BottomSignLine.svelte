<script lang="ts">
  import KeepMeSignedInBlock from '@/pages/sign-in/components/KeepMeSignedInBlock/KeepMeSignedInBlock.svelte'
  import BottomSignLineLayout from '@/pages/sign-in/components/BottomSignLine/BottomSignLineLayout.svelte'
  import ForgotPasswordBlock from '@/pages/sign-in/components/ForgotPasswordBlock/ForgotPasswordBlock.svelte'

  type BottomSignLineProps = {
    keepSignedIn: boolean
  }

  let { keepSignedIn = $bindable() }: BottomSignLineProps = $props()
</script>

<BottomSignLineLayout>
  <KeepMeSignedInBlock bind:keepSignedIn/>
  <ForgotPasswordBlock />
</BottomSignLineLayout>
