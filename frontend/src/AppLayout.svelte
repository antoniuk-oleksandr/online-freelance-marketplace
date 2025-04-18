<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { useLocation } from 'svelte-routing'
  import Header from '@/common-components/Header/Header.svelte'
  import Footer from '@/common-components/Footer/Footer.svelte'
  import PageLayout from '@/common-components/PageLayout.svelte'
  import ToastElement from '@/common-components/ToastElement/ToastElement.svelte'
  import Modal from '@/common-components/Modal/Modal.svelte'
  import ErrorComponent from '@/common-components/ErrorComponent/ErrorComponent.svelte'
  import { errorStore } from '@/common-stores/error-store'

  let { children }: LayoutProps = $props()

  let location = useLocation()

  const pagesToHideHeaderFooter = [
    'sign-in',
    'sign-up',
    'confirm-email',
    'forgot-password',
    'forgot-password',
    'reset-password',
  ]
  let hideHeaderFooter = $state<boolean>(
    pagesToHideHeaderFooter.includes($location.pathname.split('/')[1]),
  )

  const pagesToHideFooter = ['search']
  let hideFooter = $state<boolean>(pagesToHideFooter.includes($location.pathname.split('/')[1]))

  $effect(() => {
    hideHeaderFooter = pagesToHideHeaderFooter.includes($location.pathname.split('/')[1])
    hideFooter = pagesToHideFooter.includes($location.pathname.split('/')[1])

    errorStore.set({ shown: false, error: undefined })
  })
</script>

<div class="flex flex-col gap-y-6 animate-fade-in">
  <ErrorComponent />
  <Modal />
  <ToastElement />
  {#if !hideHeaderFooter}
    <Header />
  {/if}
  <PageLayout>
    <div class="min-h-app w-full">
      {@render children()}
    </div>
  </PageLayout>
  {#if !hideHeaderFooter && !hideFooter}
    <Footer />
  {/if}
</div>
