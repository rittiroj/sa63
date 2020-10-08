/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntRequisition,
    EntRequisitionFromJSON,
    EntRequisitionFromJSONTyped,
    EntRequisitionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRegisterStoreEdges
 */
export interface EntRegisterStoreEdges {
    /**
     * Requisitions holds the value of the requisitions edge.
     * @type {Array<EntRequisition>}
     * @memberof EntRegisterStoreEdges
     */
    requisitions?: Array<EntRequisition>;
}

export function EntRegisterStoreEdgesFromJSON(json: any): EntRegisterStoreEdges {
    return EntRegisterStoreEdgesFromJSONTyped(json, false);
}

export function EntRegisterStoreEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRegisterStoreEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'requisitions': !exists(json, 'requisitions') ? undefined : ((json['requisitions'] as Array<any>).map(EntRequisitionFromJSON)),
    };
}

export function EntRegisterStoreEdgesToJSON(value?: EntRegisterStoreEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'requisitions': value.requisitions === undefined ? undefined : ((value.requisitions as Array<any>).map(EntRequisitionToJSON)),
    };
}


