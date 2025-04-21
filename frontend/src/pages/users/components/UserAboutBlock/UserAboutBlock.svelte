<script lang="ts">
  import UserComponentLayout from '@/pages/users/components/UserComponentLayout.svelte'
  import Spoiler from '@/common-components/Spoiler/Spoiler.svelte'

  type UserAboutBlockProps = {
    about: string | null
  }

  const { about }: UserAboutBlockProps = $props()
</script>

{#if about !== null}
  <UserComponentLayout>
    <h2 class="text-xl font-bold">About</h2>
    <Spoiler>
      <div
        class="tiptap text-justify sm:text-pretty text-base h-fit text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
      >
        {@html about}
      </div>
    </Spoiler>
  </UserComponentLayout>
{/if}
