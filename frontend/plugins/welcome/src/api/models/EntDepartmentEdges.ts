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
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDepartmentEdges
 */
export interface EntDepartmentEdges {
    /**
     * DepartmentUser holds the value of the DepartmentUser edge.
     * @type {Array<EntUser>}
     * @memberof EntDepartmentEdges
     */
    departmentUser?: Array<EntUser>;
}

export function EntDepartmentEdgesFromJSON(json: any): EntDepartmentEdges {
    return EntDepartmentEdgesFromJSONTyped(json, false);
}

export function EntDepartmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDepartmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'departmentUser': !exists(json, 'departmentUser') ? undefined : ((json['departmentUser'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntDepartmentEdgesToJSON(value?: EntDepartmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'departmentUser': value.departmentUser === undefined ? undefined : ((value.departmentUser as Array<any>).map(EntUserToJSON)),
    };
}


