# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class AdminNodeExecutionMetaData(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'retry_group': 'str',
        'is_parent_node': 'bool',
        'spec_node_id': 'str',
        'is_dynamic': 'bool',
        'is_array': 'bool'
    }

    attribute_map = {
        'retry_group': 'retry_group',
        'is_parent_node': 'is_parent_node',
        'spec_node_id': 'spec_node_id',
        'is_dynamic': 'is_dynamic',
        'is_array': 'is_array'
    }

    def __init__(self, retry_group=None, is_parent_node=None, spec_node_id=None, is_dynamic=None, is_array=None):  # noqa: E501
        """AdminNodeExecutionMetaData - a model defined in Swagger"""  # noqa: E501

        self._retry_group = None
        self._is_parent_node = None
        self._spec_node_id = None
        self._is_dynamic = None
        self._is_array = None
        self.discriminator = None

        if retry_group is not None:
            self.retry_group = retry_group
        if is_parent_node is not None:
            self.is_parent_node = is_parent_node
        if spec_node_id is not None:
            self.spec_node_id = spec_node_id
        if is_dynamic is not None:
            self.is_dynamic = is_dynamic
        if is_array is not None:
            self.is_array = is_array

    @property
    def retry_group(self):
        """Gets the retry_group of this AdminNodeExecutionMetaData.  # noqa: E501

        Node executions are grouped depending on retries of the parent Retry group is unique within the context of a parent node.  # noqa: E501

        :return: The retry_group of this AdminNodeExecutionMetaData.  # noqa: E501
        :rtype: str
        """
        return self._retry_group

    @retry_group.setter
    def retry_group(self, retry_group):
        """Sets the retry_group of this AdminNodeExecutionMetaData.

        Node executions are grouped depending on retries of the parent Retry group is unique within the context of a parent node.  # noqa: E501

        :param retry_group: The retry_group of this AdminNodeExecutionMetaData.  # noqa: E501
        :type: str
        """

        self._retry_group = retry_group

    @property
    def is_parent_node(self):
        """Gets the is_parent_node of this AdminNodeExecutionMetaData.  # noqa: E501

        Boolean flag indicating if the node has child nodes under it This can be true when a node contains a dynamic workflow which then produces child nodes.  # noqa: E501

        :return: The is_parent_node of this AdminNodeExecutionMetaData.  # noqa: E501
        :rtype: bool
        """
        return self._is_parent_node

    @is_parent_node.setter
    def is_parent_node(self, is_parent_node):
        """Sets the is_parent_node of this AdminNodeExecutionMetaData.

        Boolean flag indicating if the node has child nodes under it This can be true when a node contains a dynamic workflow which then produces child nodes.  # noqa: E501

        :param is_parent_node: The is_parent_node of this AdminNodeExecutionMetaData.  # noqa: E501
        :type: bool
        """

        self._is_parent_node = is_parent_node

    @property
    def spec_node_id(self):
        """Gets the spec_node_id of this AdminNodeExecutionMetaData.  # noqa: E501


        :return: The spec_node_id of this AdminNodeExecutionMetaData.  # noqa: E501
        :rtype: str
        """
        return self._spec_node_id

    @spec_node_id.setter
    def spec_node_id(self, spec_node_id):
        """Sets the spec_node_id of this AdminNodeExecutionMetaData.


        :param spec_node_id: The spec_node_id of this AdminNodeExecutionMetaData.  # noqa: E501
        :type: str
        """

        self._spec_node_id = spec_node_id

    @property
    def is_dynamic(self):
        """Gets the is_dynamic of this AdminNodeExecutionMetaData.  # noqa: E501

        Boolean flag indicating if the node has contains a dynamic workflow which then produces child nodes. This is to distinguish between subworkflows and dynamic workflows which can both have is_parent_node as true.  # noqa: E501

        :return: The is_dynamic of this AdminNodeExecutionMetaData.  # noqa: E501
        :rtype: bool
        """
        return self._is_dynamic

    @is_dynamic.setter
    def is_dynamic(self, is_dynamic):
        """Sets the is_dynamic of this AdminNodeExecutionMetaData.

        Boolean flag indicating if the node has contains a dynamic workflow which then produces child nodes. This is to distinguish between subworkflows and dynamic workflows which can both have is_parent_node as true.  # noqa: E501

        :param is_dynamic: The is_dynamic of this AdminNodeExecutionMetaData.  # noqa: E501
        :type: bool
        """

        self._is_dynamic = is_dynamic

    @property
    def is_array(self):
        """Gets the is_array of this AdminNodeExecutionMetaData.  # noqa: E501

        Boolean flag indicating if the node is an array node. This is intended to uniquely identify array nodes from other nodes which can have is_parent_node as true.  # noqa: E501

        :return: The is_array of this AdminNodeExecutionMetaData.  # noqa: E501
        :rtype: bool
        """
        return self._is_array

    @is_array.setter
    def is_array(self, is_array):
        """Sets the is_array of this AdminNodeExecutionMetaData.

        Boolean flag indicating if the node is an array node. This is intended to uniquely identify array nodes from other nodes which can have is_parent_node as true.  # noqa: E501

        :param is_array: The is_array of this AdminNodeExecutionMetaData.  # noqa: E501
        :type: bool
        """

        self._is_array = is_array

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminNodeExecutionMetaData, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminNodeExecutionMetaData):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
