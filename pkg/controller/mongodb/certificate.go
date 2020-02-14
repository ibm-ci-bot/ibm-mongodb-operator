//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package mongodb

const certYaml = `
apiVersion: v1
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZ1ekNDQTZPZ0F3SUJBZ0lVUWYyM25mOXhTUmJVQ1RNSDMvZEtzbkhucGt3d0RRWUpLb1pJaHZjTkFRRUwKQlFBd2JERUxNQWtHQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdNQ2tOaGJHbG1iM0p1YVdFeEZEQVNCZ05WQkFjTQpDMHh2Y3lCQmJtZGxiR1Z6TVJnd0ZnWURWUVFLREE5WmIzVnlJRVJ2YldGcGJpQkpibU14R0RBV0JnTlZCQU1NCkQzbHZkWEl0Wkc5dFlXbHVMbU52YlRBZ0Z3MHlNREF5TVRRd09ERTJORGhhR0E4eU1USXdNREV5TVRBNE1UWTAKT0Zvd2JERUxNQWtHQTFVRUJoTUNWVk14RXpBUkJnTlZCQWdNQ2tOaGJHbG1iM0p1YVdFeEZEQVNCZ05WQkFjTQpDMHh2Y3lCQmJtZGxiR1Z6TVJnd0ZnWURWUVFLREE5WmIzVnlJRVJ2YldGcGJpQkpibU14R0RBV0JnTlZCQU1NCkQzbHZkWEl0Wkc5dFlXbHVMbU52YlRDQ0FpSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnSVBBRENDQWdvQ2dnSUIKQUt0cGtBY2pia1UvOGdMRmJNT1NhMENQTHJacTNiVnNpZ3Zvd0YrQzFUMVNOekhIQlY3Q2U4UkloYlltQTcyMQowSHpOVUJiSDRaMjZHRHJRMVVOeWNhZVp4Z2xRMW0ydmNlOGkvM1JqTFlQTGFCcnV0OXVQQzY3aEVPSTNhNVVQCkFkUXI2UjZHR250bEYxbjZmVU5KUkFYTFlqcHZDODNOcFd3VzhuSEVBTFlmM01NSndZZlV6WHBGdTlVTWVoVmcKcUpaYzBLdXJhS1BSTnNVVklDL1czOWw0R3I2VGpkSmNRUmJyOTl0WFh1cldTcHpZMXVxSDNoalFNelFSc1lUOQpEVytLaE1kelVYT1JWZFVWSEFEWGM5cHpzc3IvM1FRM1FvaWcrQmxoUHE0VXJTNGFxMnp1RStwdjRScGtEa3dKCmsyU1Jnc2ZlRCtydmpPemxXUHhVK3RtVzhQWmpmOUNPYTM4TjkrQW1LWDhLVnQ0cVU3NWpmcWZPRWhpR0VZQUIKQWFsUiszSkF6cDMwNnBEcjA2ZDRNeGsrRDRqYVRuUHBIT3h1OHFiYXFqNTYvVkVMR2xWZXJ4OVEwczZ6VUhQYQpLNzFkUmtoMnNvYmdCTDk5N3Z4SlJ2amRraFF6M1RZQ292SHVaSTR5VVFMZ1cvcWJaM1VBQTJkc2xPa3NYd3h6CjczMzhYR1llZnAzY3J6UGhiTEdQa3dlZUhsbFFQOEFkeXlrL21TOTFVWVpvVTVQN0ZtWjhWMVZObTBYOGVIdy8KQXJPQ0tGN1pCWEI1Qll6SU8yUkVCY2I3bHpVRG1QWW5qUGR1Nmc4elArUXpONCtEZ2hZdG5OclphazRTdllXaApxQWNsd3doQkFTYW80N05wd2szNllTaUhMK3FTU1VXUGJqTndNdHhoQ09vTEFnTUJBQUdqVXpCUk1CMEdBMVVkCkRnUVdCQlNMN0RoOGE4Y2V2Rkc5d1FiYXE2MGRTbDNVcURBZkJnTlZIU01FR0RBV2dCU0w3RGg4YThjZXZGRzkKd1FiYXE2MGRTbDNVcURBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElDQVFDSwpDTDg4aHNGbjkrNHg4Z3hRR2srWExtdG56bDhJRFl2N0YrVUwwTmpMNlNoT203WEVKL3ZKM252WmlDaUhxejRJCmZMNDZMb0lZbFAvNlhEbGtRendWRmNPRERIMmF0UDdjRUtpbE1BM0VoZDF5MkF4ZGlnOTQxYVBSRk1ReXBzZEkKNkNxa0dHTi9QV2h6ZE1uaFBHWnNPNmM3QUd5Wi9TYU9oMWtXM3FFRWMzRHZVd3JRSjlab1hZYzd0cGdZbU5PRQoxdEVFRkJzeERjL1dEZ1JidVp0K1lkZHlNN0s3Q2h1TytZcFljazZTTEhRQUFQR3ExUFdqUXVtUDI4SmRlZDlVCmdxbW1hUTlueTUzRDA1SDdQaE1HenZWVitLV0EyZHRaU2hMVkNZckVMNnp1OVZpdit3aTBGbDFjR2JJdXlqZUwKOEFKdWZvd3VXaFdFdW92Y0h5OHdZT1VmcmI2d0hVQnZMT0w3QmpJVUpUZ1VUQ2JQelR6eHJ2T3B3NDBCSi8zRgpUSFErb1AvMFl5aWtaeXBFczl0RXo0L0V1bzEveU1SMkNUblFCNlN0TFRwSENsdHU2SkcyYXFrdjlBZ3NHTkl2CllNRDd6ditJOTVHbVBFYzJuOVQwb2d4eVBKRnR2Rm85djdTVzZKWEMrOVJZMHkxOUpYakxLYklLT0p6dkRXWjYKWHJyUW5PKzUwS2EwZk1XQi9COGtIL29KdFN6akxYcHFTZEljN29sZUlETTNvUCtaaGk5SWxicGZjVVdwWEhlVgo4aTRMVTl1NVFyc1JwZUUzMDFZMlBDK05iZmN1RlZUOXI3aWxmcHh1VkpIb2hudDdKM2tON3FhYTJ3VEp1elE1CjNzYW44bXp2RUUwV0NpZDRYdktWZVFHdmVvNm5ScXY0cWdSbTdONmVNdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  tls.key: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUpSUUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQ1M4d2dna3JBZ0VBQW9JQ0FRQ3JhWkFISTI1RlAvSUMKeFd6RGttdEFqeTYyYXQyMWJJb0w2TUJmZ3RVOVVqY3h4d1Zld252RVNJVzJKZ085dGRCOHpWQVd4K0dkdWhnNgowTlZEY25Hbm1jWUpVTlp0cjNIdkl2OTBZeTJEeTJnYTdyZmJqd3V1NFJEaU4ydVZEd0hVSytrZWhocDdaUmRaCituMURTVVFGeTJJNmJ3dk56YVZzRnZKeHhBQzJIOXpEQ2NHSDFNMTZSYnZWREhvVllLaVdYTkNycTJpajBUYkYKRlNBdjF0L1plQnErazQzU1hFRVc2L2ZiVjE3cTFrcWMyTmJxaDk0WTBETTBFYkdFL1ExdmlvVEhjMUZ6a1ZYVgpGUndBMTNQYWM3TEsvOTBFTjBLSW9QZ1pZVDZ1RkswdUdxdHM3aFBxYitFYVpBNU1DWk5ra1lMSDNnL3E3NHpzCjVWajhWUHJabHZEMlkzL1FqbXQvRGZmZ0ppbC9DbGJlS2xPK1kzNm56aElZaGhHQUFRR3BVZnR5UU02ZDlPcVEKNjlPbmVETVpQZytJMms1ejZSenNidkttMnFvK2V2MVJDeHBWWHE4ZlVOTE9zMUJ6Mml1OVhVWklkcktHNEFTLwpmZTc4U1ViNDNaSVVNOTAyQXFMeDdtU09NbEVDNEZ2Nm0yZDFBQU5uYkpUcExGOE1jKzk5L0Z4bUhuNmQzSzh6CjRXeXhqNU1Ibmg1WlVEL0FIY3NwUDVrdmRWR0dhRk9UK3habWZGZFZUWnRGL0hoOFB3S3pnaWhlMlFWd2VRV00KeUR0a1JBWEcrNWMxQTVqMko0ejNidW9QTXova016ZVBnNElXTFp6YTJXcE9FcjJGb2FnSEpjTUlRUUVtcU9PegphY0pOK21Fb2h5L3Fra2xGajI0emNETGNZUWpxQ3dJREFRQUJBb0lDQVFDZU01RndaU3p5ME8vRnl2eDFDLy9jCjE4bGZKd1owMlRZWGc3dDQ2MEZ2bFIvSDIwMkRoYS9CR3NYOXROd2pEVmpjdG9sQ1hUeEgzR2RjY1QxRFFLN1EKNXNXMTkxdlFZK3Fkc1FsZ0crRDRFSldHdGVnT3hka2NrRENLK0dkY2JUdGMzU0lNdit4VkVwb1l1MjA4MnJQagpQVER0aVd1MWlDVXl1NE1McU1DWFZwVGpsVzkrczlMcEZNOGs5SjBBZWZOZXcwNzhBZmxUcUZZS1UxTWZsQXVKCnBLWGFXa2k4YTlUUEdQM3EvSnBuMTF1Sm96L0h1Wm9zQUl3aGRHc0l1SnpiNDFaazFuUElXOHlDR3pkN04xNS8Kc2VURUlwMlR0SVM5b1FNYU9Yd21Jb01yV0FYN1JVZ3dTUTNTMDZHMlgrZ2ZHNFUrZmg3bTBDeWF2NW4vV1ZjaQpVNE13ZldQWHJWZnNGNWhvWW5tQThYSUsvZWFja3FCaWZvL29CRTFDWGk3cGh5K2FYbmhwMlBTSXpSWXEwbVF4Ckc3WWxNb1g5QlJlT1FqaStyWXZrdmExZTRLM0Q1UmhYTUFTWXY1SHJSZS9HZ3ljQVhOQ1MxWDhXSGdHRjllWlkKdlRqaHlnMTRGYWRRdmNLOHQ3eTl3UEI5TzZhTnhzeGY4ak41dWRnUjY3RmkzTyt3L0RUNUZwNlFDYlpMWEhNNgpqWHpYc3JqMnh0Tnl1K1hNeGZvOVJhT2RFTWp0UFRNVlNEOWNiZE0xMm5iRHNhL0xCRUs2QzNVczZVcktmRFJRCnlwOUpSc3lkcGR3VFhXUUVsVjE4ZFpyc0g5QndaUTc1RStXY1hUekc2YmFlZkJQaTUxSVU1S0xCMkhtcG5BVFkKdmsxSkVHZUQ1V2JaaXBmaTJEMUVVUUtDQVFFQTQ0MHJyUFR6VThYa1ZTNDhqWjNKNFVrYytxOGxLVlRBNXdyVApyck4yVXNKNGYyRC9qYlpzaVZWakZlck85eS9KdUQyb2xRWVZBTStaSWJyTXJ0aUw0b2dyeGw1OGdzdXBDUWFXCjVZUzJwMVVHTkFNRlErM0VqcWo0WU5JcHUwcVJKZ3Y0a1FiTnVTb2ZTVVppalpxRFE3SmZhVXY0VCt1SHVPcjgKd1M3ejB0QXBXUGNFWXM4UHZEZm5SK3NjampJeXdJMWtVbmpBMTAzdjhQeG93VVB3bW5PTUI3Yk5SVExzTkxIcgp6V280bi9jQ0l3dk4yNmlmWmlDbXRyODB4WnlXZmxpRGlJZUxBdGVnenF3TmVscXRjcEtKdmp0Z1hlb2g0bEo0CmRBNTRyVS84aHZ6dndLOTFxRVB4cXdjaWpiUEczTzJMVlZlZW9VNWFnaGkzVXlhYUR3S0NBUUVBd05lbUxUbTkKNVYyMC81NWtneiswZUgvM3lMcmdOUitDelhSaXcycy9uSzk4VFhieTNVVnBGbFVhRGhZTVI5UkIwWlVwc3gzVQpLQ1k4NDB0VmxoVmE5aXVyUmljMERnSG1iK2ViVXM3enZOb3plSW1RZjlZYTE1RllMOUFTazN5Y0hab2p5S0NjCitiWURVa0lwSUpzRCtoWlBZRThaWS9HTStTTDJKL1A5SjF3UUg0RU5yYTAvVGZ1SWRoaTl5UlhGS1FwcHRyOGsKM0J5SVhNRTQvdDV1WGdzV25sSElyVVJIUWd3dWs2Y04rUk1IK1FEVk10WHBrZS9La3AvM2tCY2xDRTNFKzdMNQpwM3RTK3RvSnRRYlVCZmNaSEFkOFo4ZGlhdkJqTk5vNFpNSU5yRXlIYlV0ckNjQWtGNHVxUEkvRXZZWmpuWnZ4CjdKU21FM3dtOTN4Y1JRS0NBUUVBdjgxNXVCTE1tNXRkaUhNdHRVMEFJclluQ0NMamx3QUtqWFR4MHZzUVRGZ0IKaTVUdU03eEZwa2pqK1ZCdjNhbFpSY1FyT2xuakVFanNVYU1MN29FMi85NDV6NzVMQTFDNWpaTVJWQ3RXYnB0YgpYRElEbVN0c2w3ZFRqUUQ1RGhhZlhFdVEzR1c0dUZBS0NSQ2I5N0g2M1BBWVBrQTc4Tm1lRUQ5Nmh6ZVZhK2ZzClFrSU1YMkViYVZKUno2d2M2THJsVHFxZm1taVFXNnNsQ2FNb2hXZElBbGhNK3grWnpLOU1yR29nVVlnc2JLR2YKR24xVHFzaHFlNElnSkQvdE1uaDMzYzJzS1VQZTU2bHJzV2tTLzBRNW1TNXRqbnlzRU4xT2JGK0ppb0c4TkpJTQpVSmdaV3ZoS20yZVI1OUJlenNSdlpqK1FSNEpRZ3JUeDAyRVQvYkcyL3dLQ0FRRUFvVWY2WUZxcGVob01GYmxJCkFRa1hpNndpOU10dVZwK2JDdS9xNWR0ZG9WS1hHRWFDU2hNU0Y2NW5FeDdYZUgrSjZKbmkvVDNXUVJueDFIc2UKQWw2cjY3U01FeWtZZlJxVHVrV3UvdXU4cDlhcU00YUJ4cGl2YkZqUHMzVkhBL2kwcklGTXFpL0IrWXFEYnBtTQp3Z3REb1RabGZudkpGWEJnNGNDZXRMMTNuU0ovRFlvbWdYbWl0enJHWDFiRDYxaDh3NTFFMnFVMFU2NjlVUzM4CjJCUDVwRWVMekM5RU1iMG9CNUllcGppU0ZCMFpqdVJtYXBURE90MmN3MUpzZnlNTVFzeWdFWTRYQkQ2OUlMV1kKWnNGbU0vTDhHaGdjc01MdGlkaXdiL1NPTTU2a1J0VGVjc2NmcDFEK0huOGRpbmNIS1NjclN4YXdsNWVlNmk3cwpmQkdUTFFLQ0FRRUFpTXIvRjEyRzN3L3c0aEd1NS9vQ0o3SVZOSDJodFZDWUFyRzFjSlovWGptTGdJYlhwRTdlCk91V1VjOGl6SVBZVHUwS3dXbGZabk55RDZLQUV4R0tyWW5zYmtqTGttMVNJYXpnbHF4ZG9mSGRUelNsRjM2eGkKVnZtY0lQcnVOS0gwMVl1L2NZSVVmak1pcjRVME9iMGJ3SnkrUjI0UnZKTmt4NW04VW52TVYweVUrRjRXK3dGMQpUeUZyUklGQ1JJQVZWMGVKUm5XTzRxbTE0aXZtYlM2K0FLTmZzTlRoSHpIVFBWQzNLUTJqazdiemZmWXB3dGQxCk10ekduV0dsa01kaG41cmhVWkVUYVdxd3RpV215MVZjNUQzNlBPNUR2NHVXMW5PR29wRWJWNVlod2xxVkpXb0UKanNTemRzcDRjKzlhd2kyZ2lzaXYyMmd1TWRZUUxORmpwdz09Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
kind: Secret
metadata:
  name: cluster-ca-cert
  namespace: ibm-mongodb-operator
type: kubernetes.io/tls
`
