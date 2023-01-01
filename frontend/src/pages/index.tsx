import React from "react"
import {
    Center,
    Text,
    Box,
    Tabs,
    TabList,
    TabPanels,
    Tab,
    TabPanel,
    Accordion,
    AccordionItem,
    AccordionButton,
    AccordionPanel,
    AccordionIcon,
} from "@chakra-ui/react"

import { InferGetServerSidePropsType } from "next"

interface statusMetadata {
    reachable: boolean
    error: string
    date: string
}
interface Domain {
    configuration: {
        name: string
        domain: string
        updateInterval: number
        requirements: {
            statusCode: string
            contentType: string
        }
    }
    history: statusMetadata[]
    lastStatus: statusMetadata
}

interface StatusGroup {
    name: string
    domains: Domain[]
}

export async function getServerSideProps() {
    let statusGroups: StatusGroup[] = []

    const baseurl = process.env.BASE_URL || "http://localhost:8080"
    const requestData = await fetch(`${baseurl}/routes`)
    const jsonBody = await requestData.json()

    const proms = jsonBody.providedRoutes.map(async (route: string) => {
        if (route.startsWith("/")) route = route.substring(1, route.length)
        const statusGroupRequest = await fetch(`${baseurl}/${route}`)
        if (!statusGroupRequest.ok) {
            return {
                props: {},
            }
        }

        const statusGroupBody = await statusGroupRequest.json()
        console.log(statusGroupBody)

        const statusGroup: StatusGroup = {
            name: statusGroupBody.name,
            domains: statusGroupBody.domains,
        }
        statusGroups.push(statusGroup)
    })

    await Promise.all(proms)

    return {
        props: {
            statusgroups: statusGroups,
        },
    }
}

export default function Index({ statusgroups }: InferGetServerSidePropsType<typeof getServerSideProps>) {
    return (
        <>
            <Center padding="5em">
                <Text fontSize="3xl" fontWeight="700">
                    TestausUptime
                </Text>
            </Center>

            <Center width="100%">
                <Tabs variant="line" width="100%" maxWidth="800px" padding="0 2em">
                    <TabList>
                        {statusgroups.map((statusgroup, id) => (
                            <Tab key={id}>{statusgroup.name}</Tab>
                        ))}
                    </TabList>

                    <TabPanels>
                        {statusgroups.map((statusgroup, id) => (
                            <TabPanel key={id}>
                                <Accordion allowMultiple width="100%">
                                    {statusgroup.domains?.map((domain, domainIndex) => (
                                        <AccordionItem key={domainIndex}>
                                            <AccordionButton>
                                                {domain.lastStatus.error != "" ? (
                                                    domain.lastStatus.reachable ? (
                                                        <i className="bi bi-exclamation-circle"></i>
                                                    ) : (
                                                        <i className="bi bi-exclamation-triangle-fill"></i>
                                                    )
                                                ) : (
                                                    <i className="bi bi-check2-circle"></i>
                                                )}
                                                <Box marginLeft="1em" as="span" flex="1" textAlign="left">
                                                    {domain.configuration.domain}
                                                </Box>
                                                <AccordionIcon />
                                            </AccordionButton>
                                            <AccordionPanel pb={4}>
                                                {domain.lastStatus.error != "" ? (
                                                    domain.lastStatus.reachable ? (
                                                        <>
                                                            {domain.lastStatus.error}, but service is
                                                            reachable
                                                        </>
                                                    ) : (
                                                        <>Service is completely down.</>
                                                    )
                                                ) : (
                                                    <>Service is up and isn&apos;t experiencing any issues.</>
                                                )}
                                            </AccordionPanel>
                                        </AccordionItem>
                                    ))}
                                </Accordion>
                            </TabPanel>
                        ))}
                    </TabPanels>
                </Tabs>
            </Center>
        </>
    )
}
