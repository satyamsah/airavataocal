/*
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
*/
package org.apache.airavata.persistance.registry.jpa.resources;

import org.apache.airavata.persistance.registry.jpa.Resource;
import org.apache.airavata.persistance.registry.jpa.ResourceType;
import org.apache.airavata.persistance.registry.jpa.ResourceUtils;
import org.apache.airavata.persistance.registry.jpa.model.*;
import org.apache.airavata.persistance.registry.jpa.utils.QueryGenerator;

import javax.persistence.EntityManager;
import javax.persistence.Query;
import java.util.ArrayList;
import java.util.List;

public class ApplicationDescriptorResource extends AbstractResource {
    private String name;
    private String gatewayName;
    private String updatedUser;
    private String content;
    private String hostDescName;
    private String serviceDescName;

    /**
     *
     * @param name application descriptor name
     * @param gatewayName  gateway name
     * @param hostDescName host descriptor name
     * @param serviceDescName service descriptor name
     * returns ApplicationDescriptorResource
     */
    public ApplicationDescriptorResource(String name, String gatewayName,
                                         String hostDescName, String serviceDescName) {
        this.setName(name);
        this.gatewayName = gatewayName;
        this.hostDescName = hostDescName;
        this.serviceDescName = serviceDescName;
    }

    /**
     *
     */
    public ApplicationDescriptorResource() {
    }

    /**
     *
     * @return gateway name
     */
    public String getGatewayName() {
        return gatewayName;
    }

    /**
     *
     * @param gatewayName gateway name
     */
    public void setGatewayName(String gatewayName) {
        this.gatewayName = gatewayName;
    }

    /**
     *
     * @param updatedUser updated user
     */
    public void setUpdatedUser(String updatedUser) {
        this.updatedUser = updatedUser;
    }

    /**
     *
     * @return name of the application descriptor
     */
    public String getName() {
        return name;
    }

    /**
     *
     * @return content
     */
    public String getContent() {
        return content;
    }

    /**
     *
     * @return  host descriptor name
     */
    public String getHostDescName() {
        return hostDescName;
    }

    /**
     *
     * @return service descriptor name
     */
    public String getServiceDescName() {
        return serviceDescName;
    }

    /**
     *
     * @param content content of the application descriptor
     */
    public void setContent(String content) {
        this.content = content;
    }

    /**
     *
     * @param hostDescName host descriptor name
     */
    public void setHostDescName(String hostDescName) {
        this.hostDescName = hostDescName;
    }

    /**
     *
     * @param serviceDescName  service descriptor name
     */
    public void setServiceDescName(String serviceDescName) {
        this.serviceDescName = serviceDescName;
    }

    /**
     * Since application descriptors are at the leaf level, this method is not
     * valid for application descriptors
     * @param type  child resource types
     * @return UnsupportedOperationException
     */
    public Resource create(ResourceType type) {
        throw new UnsupportedOperationException();
    }

    /**
     * Since application descriptors are at the leaf level, this method is not
     * valid for application descriptors
     * @param type child resource types
     * @param name name of the resource
     */
    public void remove(ResourceType type, Object name) {
        throw new UnsupportedOperationException();
    }

    /**
     * key should be gateway_name, application_descriptor_ID, host_descriptor_ID, service_descriptor_ID
     *
     * @param keys primary keys of the Application_descriptor table
     */
//    public void removeMe(Object[] keys) {
//        EntityManager em = ResourceUtils.getEntityManager();
//        em.getTransaction().begin();
//        QueryGenerator queryGenerator = new QueryGenerator(APPLICATION_DESCRIPTOR);
//        queryGenerator.setParameter(ApplicationDescriptorConstants.GATEWAY_NAME, keys[0]);
//        queryGenerator.setParameter(ApplicationDescriptorConstants.APPLICATION_DESC_ID, keys[1]);
//        queryGenerator.setParameter(ApplicationDescriptorConstants.HOST_DESC_ID, keys[2]);
//        queryGenerator.setParameter(ApplicationDescriptorConstants.SERVICE_DESC_ID, keys[3]);
//        Query q = queryGenerator.deleteQuery(em);
//        q.executeUpdate();
//        em.getTransaction().commit();
//        em.close();
//    }

    /**
     *
     * Since application descriptors are at the leaf level, this method is not
     * valid for application descriptors
     * @param type child resource types
     * @param name name of the resource
     * @return UnsupportedOperationException
     */
    public Resource get(ResourceType type, Object name) {
        throw new UnsupportedOperationException();
    }

    /**
     * keys should contain gateway_name, application_descriptor_ID, host_descriptor_ID, service_descriptor_ID
     *
     * @param keys names
     * @return list of ApplicationDescriptorResources
     */
    public List<Resource> populate(Object[] keys) {
        List<Resource> list = new ArrayList<Resource>();
        EntityManager em = ResourceUtils.getEntityManager();
        em.getTransaction().begin();
        QueryGenerator queryGenerator = new QueryGenerator(APPLICATION_DESCRIPTOR);
        queryGenerator.setParameter(ApplicationDescriptorConstants.GATEWAY_NAME, keys[0]);
        queryGenerator.setParameter(ApplicationDescriptorConstants.APPLICATION_DESC_ID, keys[1]);
        queryGenerator.setParameter(ApplicationDescriptorConstants.HOST_DESC_ID, keys[2]);
        queryGenerator.setParameter(ApplicationDescriptorConstants.SERVICE_DESC_ID, keys[3]);
        Query q = queryGenerator.selectQuery(em);
        Application_Descriptor applicationDescriptor = (Application_Descriptor) q.getSingleResult();
        ApplicationDescriptorResource applicationDescriptorResource =
                (ApplicationDescriptorResource) Utils.getResource(
                        ResourceType.APPLICATION_DESCRIPTOR, applicationDescriptor);
        em.getTransaction().commit();
        em.close();
        list.add(applicationDescriptorResource);
        return list;
    }

    /**
     * Since application descriptors are at the leaf level, this method is not
     * valid for application descriptors
     * @param type child resource types
     * @return UnsupportedOperationException
     */
    public List<Resource> get(ResourceType type) {
        throw new UnsupportedOperationException();
    }

    /**
     *  save application descriptor to database
     */
    public void save() {
        EntityManager em = ResourceUtils.getEntityManager();
        Application_Descriptor existingAppDesc = em.find(Application_Descriptor.class, new Application_Descriptor_PK(gatewayName,
                name, hostDescName, serviceDescName));
        em.close();

        em = ResourceUtils.getEntityManager();
        em.getTransaction().begin();
        Application_Descriptor applicationDescriptor = new Application_Descriptor();
        applicationDescriptor.setApplication_descriptor_ID(getName());
        Gateway gateway = new Gateway();
        gateway.setGateway_name(gatewayName);
        applicationDescriptor.setGateway(gateway);
        Users user = new Users();
        user.setUser_name(updatedUser);
        applicationDescriptor.setUser(user);
        applicationDescriptor.setApplication_descriptor_xml(content);
        applicationDescriptor.setService_descriptor_ID(serviceDescName);
        applicationDescriptor.setHost_descriptor_ID(hostDescName);
        if (existingAppDesc != null) {
            applicationDescriptor.setUser(user);
            applicationDescriptor.setApplication_descriptor_xml(content);
            applicationDescriptor = em.merge(existingAppDesc);
        } else {
            em.merge(applicationDescriptor);
        }

        em.getTransaction().commit();
        em.close();
    }

    /**
     * Since application descriptors are at the leaf level, this method is not
     * valid for application descriptors
     * @param type child resource types
     * @param name name of the resource
     * @return UnsupportedOperationException
     */
    public boolean isExists(ResourceType type, Object name) {
        throw new UnsupportedOperationException();
    }

    /**
     *
     * @param name application descriptor name
     */
    public void setName(String name) {
        this.name = name;
    }
}
