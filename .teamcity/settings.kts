import jetbrains.buildServer.configs.kotlin.*
import jetbrains.buildServer.configs.kotlin.buildFeatures.perfmon
import jetbrains.buildServer.configs.kotlin.buildSteps.SSHUpload
import jetbrains.buildServer.configs.kotlin.buildSteps.script
import jetbrains.buildServer.configs.kotlin.buildSteps.sshUpload
import jetbrains.buildServer.configs.kotlin.projectFeatures.githubConnection
import jetbrains.buildServer.configs.kotlin.triggers.vcs

/*
The settings script is an entry point for defining a TeamCity
project hierarchy. The script should contain a single call to the
project() function with a Project instance or an init function as
an argument.

VcsRoots, BuildTypes, Templates, and subprojects can be
registered inside the project using the vcsRoot(), buildType(),
template(), and subProject() methods respectively.

To debug settings scripts in command-line, run the

    mvnDebug org.jetbrains.teamcity:teamcity-configs-maven-plugin:generate

command and attach your debugger to the port 8000.

To debug in IntelliJ Idea, open the 'Maven Projects' tool window (View
-> Tool Windows -> Maven Projects), find the generate task node
(Plugins -> teamcity-configs -> teamcity-configs:generate), the
'Debug' option is available in the context menu for the task.
*/

version = "2023.05"

project {

    buildType(Build)

    features {
        githubConnection {
            id = "PROJECT_EXT_2"
            displayName = "GitHub.com"
            clientId = "3d0bf7d8110d0ef94e82"
            clientSecret = "credentialsJSON:d2f068fb-831d-4294-a073-6d1cf8975344"
        }
    }
}

object Build : BuildType({
    name = "Build app"
    description = "Создание исполняемого файла"

    artifactRules = """
        t_app => t_app_%build.counter%.zip
        README.md => t_app_%build.counter%.zip
    """.trimIndent()
    publishArtifacts = PublishMode.SUCCESSFUL

    params {
        param("env.GOOS", "linux")
        param("env.GOARCH", "386")
    }

    vcs {
        root(DslContext.settingsRoot)
    }

    steps {
        script {
            name = "build go project"
            scriptContent = "go build -o t_app main.go"
        }
        sshUpload {
            name = "depoly"
            transportProtocol = SSHUpload.TransportProtocol.SCP
            sourcePath = "*.zip"
            targetUrl = "bpm.dev.itkn.ru:/home/askorohodov/project/torrent-to-url-telegram-bot"
            authMethod = defaultPrivateKey {
                username = "askorohodov"
                passphrase = "credentialsJSON:e93b13a6-1431-41a3-92df-3b065124466c"
            }
        }
    }

    triggers {
        vcs {
        }
    }

    features {
        perfmon {
        }
    }
})
