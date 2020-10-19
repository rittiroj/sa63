/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDrug,
    EntDrugFromJSON,
    EntDrugFromJSONTyped,
    EntDrugToJSON,
    EntRegisterStore,
    EntRegisterStoreFromJSON,
    EntRegisterStoreFromJSONTyped,
    EntRegisterStoreToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRequisitionEdges
 */
export interface EntRequisitionEdges {
    /**
     * 
     * @type {EntDrug}
     * @memberof EntRequisitionEdges
     */
    drug?: EntDrug;
    /**
     * 
     * @type {EntRegisterStore}
     * @memberof EntRequisitionEdges
     */
    registerstore?: EntRegisterStore;
    /**
     * 
     * @type {EntUser}
     * @memberof EntRequisitionEdges
     */
    user?: EntUser;
}

export function EntRequisitionEdgesFromJSON(json: any): EntRequisitionEdges {
    return EntRequisitionEdgesFromJSONTyped(json, false);
}

export function EntRequisitionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRequisitionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drug': !exists(json, 'drug') ? undefined : EntDrugFromJSON(json['drug']),
        'registerstore': !exists(json, 'registerstore') ? undefined : EntRegisterStoreFromJSON(json['registerstore']),
        'user': !exists(json, 'user') ? undefined : EntUserFromJSON(json['user']),
    };
}

export function EntRequisitionEdgesToJSON(value?: EntRequisitionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'drug': EntDrugToJSON(value.drug),
        'registerstore': EntRegisterStoreToJSON(value.registerstore),
        'user': EntUserToJSON(value.user),
    };
}


