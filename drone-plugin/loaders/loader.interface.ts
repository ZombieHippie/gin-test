import { ArtifactUpload } from '../lib-ts/upload/artifact-upload.model'

export interface Loader {
  (artifact: ArtifactUpload): Error
} 
export interface LoaderFactory {
  new (query: {[key: string]: string }): Loader
}
