<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import MyProfileEditPageAvatarBlockLayout from './MyProfileEditPageAvatarBlockLayout.svelte'
  import AttachFilesButton from '@/common-components/AttachFilesButton/AttachFilesButton.svelte'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { Writable } from 'svelte/store'
  import { onDestroy } from 'svelte'
  import { getFileURL } from '@/utils/file-utils'

  type MyProfileEditPageAvatarBlockProps = {
    formDataStore: Writable<MyProfileEditPageFormData>
  }

  const { formDataStore }: MyProfileEditPageAvatarBlockProps = $props()

  let avatarSrc = $state<string | undefined | null>()

  const unsubscribe = formDataStore.subscribe((value) => {
    if (value.avatar instanceof File) {
      getFileURL(value.avatar).then((url) => (avatarSrc = url))
    } else avatarSrc = value.avatar
  })

  const addFiles = (files: FileList) => {
    formDataStore.update((prev) => ({ ...prev, avatar: files[0] }))
  }

  onDestroy(() => unsubscribe())
</script>

<MyProfileEditPageAvatarBlockLayout>
  {#if avatarSrc !== undefined}
    <Avatar size="48" src={avatarSrc} alt="avatar" />
    <AttachFilesButton accept="image" icon="" text="Change Avatar" {addFiles} />
  {/if}
</MyProfileEditPageAvatarBlockLayout>
