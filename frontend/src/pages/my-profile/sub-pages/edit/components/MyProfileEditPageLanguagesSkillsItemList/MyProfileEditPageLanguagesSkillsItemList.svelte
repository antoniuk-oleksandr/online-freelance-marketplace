<script lang="ts">
  import Badge from '@/common-components/Badge/Badge.svelte'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import { onMount, onDestroy } from 'svelte'
  import type { Writable } from 'svelte/store'
  import MyProfileEditPageLanguagesSkillsItemListLayout from './MyProfileEditPageLanguagesSkillsItemListLayout.svelte'
    import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'

  type MyProfileEditPageLanguagesSkillsListProps = {
    formDataStore: Writable<MyProfileEditPageFormData>
    type: 'languages' | 'skills'
  }

  const { formDataStore, type }: MyProfileEditPageLanguagesSkillsListProps = $props()

  let isMobile = $state(window.innerWidth < 768)

  const handleResize = () => (isMobile = window.innerWidth < 768)
  onMount(() => window.addEventListener('resize', handleResize))
  onDestroy(() => window.removeEventListener('resize', handleResize))

  const deleteFunc = (index: number) => {
    formDataStore.update((prev) => ({
      ...prev,
      [type]: prev[type].filter((item) => item.id !== index),
    }))
  }
</script>

<MyProfileEditPageLanguagesSkillsItemListLayout>
  {@const data = type === 'languages' ? $formDataStore.languages : $formDataStore.skills}
  {#if data.length === 0}
    <NoContentMessage styles="w-full text-center py-[5px]" text="No {type} selected yet"/>
  {:else}
  {#each data as item}
    <Badge {deleteFunc} {isMobile} index={item.id} useDeleteButton text={item.name} />
  {/each}
  {/if}
</MyProfileEditPageLanguagesSkillsItemListLayout>
