<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import { themeStore } from '@/common-stores/theme-storage'
  import HeaderProfileBlockUnsignedBlockLayout from './HeaderProfileBlockUnsignedBlockLayout.svelte'
  import HeaderProfileBlockUnsignedBlockIcon from './components/HeaderProfileBlockUnsignedBlockIcon/HeaderProfileBlockUnsignedBlockIcon.svelte'

  type HeaderProfileBlockUnsignedBlockProps = {
    darkMode: boolean | null
  }

  const { darkMode }: HeaderProfileBlockUnsignedBlockProps = $props()
</script>

<HeaderProfileBlockUnsignedBlockLayout>
  <Button clickAction={() => themeStore.set(!darkMode)} color="outline">
    <HeaderProfileBlockUnsignedBlockIcon {darkMode} />
  </Button>
  <Button styles="w-32" link="/sign-in">Sign In</Button>
</HeaderProfileBlockUnsignedBlockLayout>
