<script lang="ts">
  import CreateService from '../../components/CreateService/CreateService.svelte'
  import CreateServiceRequirementsFormContent from './components/CreateServiceRequirementsFormContent/CreateServiceRequirementsFormContent.svelte'
  import { handleCreateServiceRequirementsFormSubmit } from './handlers'
  import { createServiceRequirementsFormSchema } from './helpers'
</script>

<CreateService
  schema={createServiceRequirementsFormSchema}
  onSubmit={handleCreateServiceRequirementsFormSubmit}
  activeStepIndex={2}
>
  <CreateServiceRequirementsFormContent />
</CreateService>
