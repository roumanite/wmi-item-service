import * as cdk from '@aws-cdk/core';
import * as codepipeline from '@aws-cdk/aws-codepipeline';
import { CodeCommitSourceAction, CodeBuildAction } from '@aws-cdk/aws-codepipeline-actions';
import * as codecommit from '@aws-cdk/aws-codecommit';
import * as codebuild from '@aws-cdk/aws-codebuild';
import * as ecr from '@aws-cdk/aws-ecr';

export class WmiItemServiceStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const ecrRepo = new ecr.Repository(this, 'WmiRepo');
/*
    const pipeline = new codepipeline.Pipeline(this, 'WmiPipeline', {
      pipelineName: 'WmiPipeline',
      crossAccountKeys: false,
    });
    
    const sourceOutput = new codepipeline.Artifact();
    const buildOutput = new codepipeline.Artifact();
    pipeline.addStage({
      stageName: "Source",
      actions: [new CodeCommitSourceAction({
        actionName: "Source-Action",
        output: sourceOutput,
        repository: new codecommit.Repository(this, 'WmiRepoResource', {
          repositoryName: 'WmiRepo',
        }),
      })],
    });

    pipeline.addStage({
      stageName: "CodeBuild_Action",
      actions: [new CodeBuildAction({
        actionName: "CodeBuild_Action",
        input: sourceOutput,
        outputs: [buildOutput],
        project: new codebuild.PipelineProject(
          this,
          'Project',
          {
          buildSpec: this.createBuildSpec(),
          environment: {
          buildImage: codebuild.LinuxBuildImage.STANDARD_2_0,
          privileged: true,
          },
          environmentVariables: {
          REPOSITORY_URI: {value: this.ecrRepo.repositoryUri},
          CONTAINER_NAME: {value: this.containerName}
          }
          }
          );
      })],
    })
*/
  }
}


/*
 private createImageBuildStage(
 stageName: string,
 input: codepipeline.Artifact,
 output: codepipeline.Artifact
 ): codepipeline.StageProps {
 const project = new codebuild.PipelineProject(
 this,
 'Project',
 {
 buildSpec: this.createBuildSpec(),
 environment: {
 buildImage: codebuild.LinuxBuildImage.STANDARD_2_0,
 privileged: true,
 },
 environmentVariables: {
 REPOSITORY_URI: {value: this.ecrRepo.repositoryUri},
 CONTAINER_NAME: {value: this.containerName}
 }
 }
 );
 this.ecrRepo.grantPullPush(project.grantPrincipal);
 
 const codebuildAction = new codepipeline_actions.CodeBuildAction({
 actionName: 'CodeBuild_Action',
 input: input,
 outputs: [output],
 project: project,
 });

 return {
 stageName: stageName,
 actions: [codebuildAction],
 };
 }
 */
