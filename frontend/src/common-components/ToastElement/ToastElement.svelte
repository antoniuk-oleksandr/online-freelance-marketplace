<script lang="ts">
  import ToastElementLayout from '@/common-components/ToastElement/ToastElementLayout.svelte'
  import { toastElementStore } from '@/common-components/ToastElement/store/toast-element-store'
  import type { ToastElementStore } from '@/types/ToastElementStore.ts'
  import ToastElementIcon from '@/common-components/ToastElement/components/ToastElementIcon/ToastElementIcon.svelte'
  import ToastElementMessage from '@/common-components/ToastElement/components/ToastElementMessage/ToastElementMessage.svelte'
  import ToastElementCloseButton from '@/common-components/ToastElement/components/ToastElementCloseButton/ToastElementCloseButton.svelte'

  let timeoutVal = $state(0)
  let toastValue = $state<ToastElementStore | undefined>()
  toastElementStore.subscribe((value) => {
    timeoutVal && clearTimeout(timeoutVal)

    if (value.show) {
      timeoutVal = setTimeout(() => {
        toastElementStore.update((prev) => ({ ...prev, show: false }))
      }, 5000)
    }

    toastValue = value
  })
</script>

{#if toastValue && toastValue.show}
  <ToastElementLayout {...toastValue}>
    <ToastElementIcon {...toastValue} />
    <ToastElementMessage {...toastValue} />
    <ToastElementCloseButton />
  </ToastElementLayout>
{/if}
