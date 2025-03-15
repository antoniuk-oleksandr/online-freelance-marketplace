<script lang="ts">
  import { onMount } from 'svelte'
  import HeaderProfileBlockSignedInBlock from './components/HeaderProfileBlockSignedInBlock/HeaderProfileBlockSignedInBlock.svelte'
  import HeaderProfileBlockUnsignedBlock from './components/HeaderProfileBlockUnsignedBlock/HeaderProfileBlockUnsignedBlock.svelte'
  import type { SignHeaderData } from '@/types/SignHeaderData'
  import { getUserSession } from '../../helpers'
  import { themeStore } from '@/common-stores/theme-storage'
  import { signDataStore } from './sign-data-store'

  let signData = $state<SignHeaderData | undefined>()
  signDataStore.subscribe((value) => (signData = value))

  let darkMode = $state<boolean | null>(null)
  themeStore.subscribe((value) => (darkMode = value))

  onMount(() => signDataStore.set(getUserSession()))
</script>

{#if signData}
  {#if signData.signedIn}
    <HeaderProfileBlockSignedInBlock {darkMode} {signData} />
  {:else}
    <HeaderProfileBlockUnsignedBlock {darkMode}/>
  {/if}
{/if}
