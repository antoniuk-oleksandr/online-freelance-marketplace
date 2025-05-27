<script lang="ts">
  import { themeStore } from '@/common-stores/theme-storage'
  import type { SignHeaderData } from '@/types/SignHeaderData'
  import { onMount } from 'svelte'
  import { getUserSession } from '../../helpers'
  import HeaderProfileBlockSignedInBlock from './components/HeaderProfileBlockSignedInBlock/HeaderProfileBlockSignedInBlock.svelte'
  import HeaderProfileBlockUnsignedBlock from './components/HeaderProfileBlockUnsignedBlock/HeaderProfileBlockUnsignedBlock.svelte'
  import { signDataStore } from './sign-data-store'

  let signData = $state<SignHeaderData | undefined>()
  signDataStore.subscribe((value) => (signData = value))

  let darkMode = $state<boolean | null>(null)
  themeStore.subscribe((value) => (darkMode = value))
</script>

{#if signData}
  {#if signData.authenticated}
    <HeaderProfileBlockSignedInBlock {darkMode} {signData} />
  {:else}
    <HeaderProfileBlockUnsignedBlock {darkMode} />
  {/if}
{/if}
